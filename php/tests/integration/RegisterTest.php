<?php

declare(strict_types=1);

namespace kuaukutsu\ec\grpc\tests\integration;

use Amp\TimeoutCancellation;
use Faker\Factory;
use Faker\Generator;
use Thesis\Grpc\GrpcException;
use PHPUnit\Framework\Attributes\CoversNothing;
use PHPUnit\Framework\TestCase;
use kuaukutsu\ec\grpc\generate\php\auth\RegisterRequest;
use kuaukutsu\ec\grpc\tests\ServiceFactory;

#[CoversNothing]
final class RegisterTest extends TestCase
{
    private static Generator $faker;

    public static function setUpBeforeClass(): void
    {
        self::$faker = Factory::create();
    }

    public function testResponseHappened(): void
    {
        $service = ServiceFactory::makeAuthService();

        $response = $service->register(
            request: new RegisterRequest(
                email: self::$faker->email(),
                password: self::$faker->password(6, 12),
            ),
            cancellation: new TimeoutCancellation(3.)
        );

        self::assertNotEmpty($response->uuid);
    }

    public function testDuplicateEmail(): void
    {
        $service = ServiceFactory::makeAuthService();

        $email = self::$faker->email();

        $service->register(
            request: new RegisterRequest(
                email: $email,
                password: self::$faker->password(6, 12),
            ),
            cancellation: new TimeoutCancellation(3.)
        );

        // A grpc error with status code "ALREADY_EXISTS" and message "user already exists" received
        self::expectException(GrpcException::class);

        $service->register(
            request: new RegisterRequest(
                email: $email,
                password: self::$faker->password(6, 12),
            ),
            cancellation: new TimeoutCancellation(3.)
        );
    }
}
