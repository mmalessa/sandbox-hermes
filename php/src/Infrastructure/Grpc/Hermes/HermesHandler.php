<?php

declare(strict_types=1);

namespace App\Infrastructure\Grpc\Hermes;

use Spiral\RoadRunner\GRPC\ContextInterface;

class HermesHandler implements HermesHandlerInterface
{
    public function Handle(ContextInterface $ctx, HermesRequest $in): HermesResponse
    {
        $response = new HermesResponse();
        $response->setResult('Hello from PHP! (received message: ' . $in->getBody() . ')');
        return $response;
    }
}
