<?php

declare(strict_types=1);

namespace kuaukutsu\ec\grpc\tests\integration;

use Amp\TimeoutCancellation;
use Auth\LoginRequest;
use Auth\RegisterRequest;
use Faker\Factory;
use Faker\Generator;
use Firebase\JWT\JWT;
use Firebase\JWT\Key;
use kuaukutsu\ec\grpc\tests\ServiceFactory;
use PHPUnit\Framework\Attributes\CoversNothing;
use PHPUnit\Framework\TestCase;

#[CoversNothing]
final class LoginTest extends TestCase
{
    private static Generator $faker;

    public static function setUpBeforeClass(): void
    {
        self::$faker = Factory::create();
    }

    public function testResponseHappened(): void
    {
        $service = ServiceFactory::makeAuthService();

        $email = self::$faker->email();
        $pass = self::$faker->password(6, 12);

        $responseRegister = $service->register(
            request: new RegisterRequest(
                email: $email,
                password: $pass,
            ),
            cancellation: new TimeoutCancellation(3.)
        );

        self::assertNotEmpty($responseRegister->uuid);

        $responseLogin = $service->login(
            request: new LoginRequest(
                email: $email,
                password: $pass,
                appId: 1,
            ),
            cancellation: new TimeoutCancellation(3.)
        );

        self::assertNotEmpty($responseLogin->token);

        $payload = JWT::decode(
            $responseLogin->token,
            new Key('718e4894-a518-4802-9205-4838c7ddbd42', 'HS256'),
        );

        self::assertEquals($email, $payload->email);
        self::assertEquals($responseRegister->uuid, $payload->uuid);
    }
}
