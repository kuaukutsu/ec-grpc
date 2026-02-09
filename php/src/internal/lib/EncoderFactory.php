<?php

declare(strict_types=1);

namespace kuaukutsu\ec\grpc\internal\lib;

use Thesis\Grpc\Encoding\Encoder;
use Thesis\Protobuf\Reflection\Reflector;
use Thesis\Protobuf\Serializer;

/**
 * @internal
 */
final readonly class EncoderFactory
{
    public static function makeProtobuf(): Encoder
    {
        return  new ProtobufEncoder(
            new Serializer(),
            Reflector::build(),
        );
    }
}
