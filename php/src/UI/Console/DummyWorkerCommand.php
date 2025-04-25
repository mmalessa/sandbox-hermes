<?php

declare(strict_types=1);

namespace App\UI\Console;

use App\Infrastructure\Grpc\Hermes\HermesHandler;
use Symfony\Component\Console\Attribute\AsCommand;
use Symfony\Component\Console\Command\Command;
use Symfony\Component\Console\Input\InputInterface;
use Symfony\Component\Console\Output\OutputInterface;

#[AsCommand(name: 'dummy:worker', description: 'gRPC Worker')]
class DummyWorkerCommand extends Command
{
    public function __construct(
        private readonly HermesHandler $hermesHandler
    )
    {
        parent::__construct();
    }

    protected function execute(InputInterface $input, OutputInterface $output): int
    {
        while (true) {}

        return Command::SUCCESS;
    }
}
