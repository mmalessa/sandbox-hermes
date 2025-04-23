<?php

declare(strict_types=1);

namespace App\UI\Console;

use App\Infrastructure\Grpc\Hermes\HermesRequest;
use Symfony\Component\Console\Attribute\AsCommand;
use Symfony\Component\Console\Command\Command;
use Symfony\Component\Console\Input\InputInterface;
use Symfony\Component\Console\Output\OutputInterface;

#[AsCommand(name: 'grpc:serialize-test', description: 'gRPC serialize test')]
class SerializeTestCommand extends Command
{
    protected function execute(InputInterface $input, OutputInterface $output): int
    {
        $request = new HermesRequest();
        $request->setPayload("Hello");
        $data = $request->serializeToString();
        echo bin2hex($data);
        return Command::SUCCESS;
    }
}
