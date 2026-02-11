<?php

declare(strict_types=1);

namespace kuaukutsu\ec\grpc\tests;

use Auth\AuthServiceClient;
use Thesis\Grpc\Client\Builder;
use kuaukutsu\ec\grpc\internal\lib\EncoderFactory;

final readonly class ServiceFactory
{
    public static function makeAuthService(): AuthServiceClient
    {
        $client = new Builder(EncoderFactory::makeProtobuf())
            ->withHost('http://host.docker.internal:3001')
            ->build();

        return new AuthServiceClient($client);
    }
}
