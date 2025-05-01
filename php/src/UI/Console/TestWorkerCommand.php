<?php

declare(strict_types=1);

namespace App\UI\Console;

use Symfony\Component\Console\Attribute\AsCommand;
use Symfony\Component\Console\Command\Command;
use Symfony\Component\Console\Input\InputInterface;
use Symfony\Component\Console\Output\OutputInterface;

#[AsCommand(name: 'hermes:test-worker', description: 'Hermes test worker')]
class TestWorkerCommand extends Command
{
    protected function execute(InputInterface $input, OutputInterface $output): int
    {
        while (($line = fgets(STDIN)) !== false) {
            $line = rtrim($line);
            $testEnv = getenv("TEST_ENV");
            printf("PHP received: %s (TEST_ENV=%s)\n", $line, $testEnv ?: "--");
            sleep(1);
        }

        return Command::SUCCESS;
    }
}
