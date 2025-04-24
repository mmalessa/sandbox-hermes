<?php

declare(strict_types=1);

namespace App\UI\Console;

use App\Infrastructure\Grpc\Hermes\HermesHandler;
use Spiral\RoadRunner\GRPC\Server;
use Symfony\Component\Console\Attribute\AsCommand;
use Symfony\Component\Console\Command\Command;
use Symfony\Component\Console\Input\InputInterface;
use Symfony\Component\Console\Output\OutputInterface;

#[AsCommand(name: 'grpc:worker', description: 'gRPC Worker')]
class GrpcWorkerCommand extends Command
{
    public function __construct(
        private readonly HermesHandler $hermesHandler
    )
    {
        parent::__construct();
    }

    protected function execute(InputInterface $input, OutputInterface $output): int
    {
        $server = new Server();
        $server->registerService(HermesHandler::class, $this->hermesHandler);
        $server->serve();

        return Command::SUCCESS;
    }
}
