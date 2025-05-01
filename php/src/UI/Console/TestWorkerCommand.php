<?php

declare(strict_types=1);

namespace App\UI\Console;

use MessagePack\MessagePack;
use Symfony\Component\Console\Attribute\AsCommand;
use Symfony\Component\Console\Command\Command;
use Symfony\Component\Console\Input\InputInterface;
use Symfony\Component\Console\Output\OutputInterface;

#[AsCommand(name: 'hermes:test-worker', description: 'Hermes test worker')]
class TestWorkerCommand extends Command
{
    protected function execute(InputInterface $input, OutputInterface $output): int
    {
        $testEnv = getenv("TEST_ENV");

        while (!feof(STDIN)) {
            $data = $this->readMessage();
            if ($data === null) {
                break;
            }
            $request = MessagePack::unpack($data);

            $responseMessage = sprintf("Response from PHP: (request %s)", json_encode($request));

            $response = [
                'status' => 'SUCCESS',
                'statusCode' => 200,
                'id' => $request['id'],
                'message' => $responseMessage,
            ];
            $this->writeMessage(MessagePack::pack($response));

            usleep(500000);
        }

        return Command::SUCCESS;
    }

    private function readMessage(): ?string
    {
        $lenBin = fread(STDIN, 4);
        if (strlen($lenBin) < 4) {
            return null;
        }
        $unpacked = unpack('Nlen', $lenBin);
        return fread(STDIN, $unpacked['len']);
    }

    function writeMessage(string $payload): void {
        $len = strlen($payload);
        fwrite(STDOUT, pack('N', $len));
        fwrite(STDOUT, $payload);
    }
}
