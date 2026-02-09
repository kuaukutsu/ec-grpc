<?php

declare(strict_types=1);

namespace kuaukutsu\ec\grpc;

use Amp\TimeoutCancellation;
use Auth\AuthServiceClient;
use Auth\RegisterRequest;
use Faker\Factory;
use Thesis\Grpc\Client\Builder;
use kuaukutsu\ec\grpc\internal\lib\EncoderFactory;

require_once __DIR__ . '/../vendor/autoload.php';

$client = new Builder(EncoderFactory::makeProtobuf())
    ->withHost('http://host.docker.internal:3001')
    ->build();

$faker = Factory::create();

$service = new AuthServiceClient($client);
$response = $service->register(
    request: new RegisterRequest(
        email: $faker->email(),
        password: $faker->password(6, 12),
    ),
    cancellation: new TimeoutCancellation(3.)
);

trap($response)->depth(2);
