<?php

declare(strict_types=1);

namespace kuaukutsu\ec\grpc\tests;

use Thesis\Grpc\Client\Builder;
use Thesis\Grpc\Protobuf\ProtobufEncoder;
use Thesis\Protobuf\Decoder;
use Thesis\Protobuf\Encoder;
use kuaukutsu\ec\grpc\generate\php\auth\AuthServiceClient;

final readonly class ServiceFactory
{
    public static function makeAuthService(): AuthServiceClient
    {
        $encoder = Encoder\Builder::buildDefault();
        $decoder = Decoder\Builder::buildDefault();

        $client = new Builder()
            ->withHost('http://host.docker.internal:3001')
            ->withProtobuf($decoder)
            ->withEncoding(new ProtobufEncoder($encoder, $decoder))
            ->build();

        return new AuthServiceClient($client);
    }
}
