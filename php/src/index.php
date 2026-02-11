<?php

declare(strict_types=1);

namespace kuaukutsu\ec\grpc;

use Amp\TimeoutCancellation;
use Auth\AuthServiceClient;
use Auth\LoginRequest;
use Auth\RegisterRequest;
use Faker\Factory;
use Firebase\JWT\JWT;
use Firebase\JWT\Key;
use Thesis\Grpc\Client\Builder;
use kuaukutsu\ec\grpc\internal\lib\EncoderFactory;

require_once __DIR__ . '/../vendor/autoload.php';

$client = new Builder(EncoderFactory::makeProtobuf())
    ->withHost('http://host.docker.internal:3001')
    ->build();

$faker = Factory::create();
$email = $faker->email();
$pass = $faker->password(6, 12);


$service = new AuthServiceClient($client);
$response = $service->register(
    request: new RegisterRequest(
        email: $email,
        password: $pass,
    ),
    cancellation: new TimeoutCancellation(3.)
);

trap($response)->depth(2);

$responseLogin = $service->login(
    request: new LoginRequest(
        email: $email,
        password: $pass,
        appId: 1,
    ),
    cancellation: new TimeoutCancellation(3.)
);

trap($responseLogin)->depth(2);

$payload = JWT::decode($responseLogin->token, new Key('718e4894-a518-4802-9205-4838c7ddbd42','HS256'));
trap($payload)->depth(2);
