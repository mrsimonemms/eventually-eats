import { Client, Connection } from '@temporalio/client';

// Handle singleton
let client: Client;

export async function ensureConnection(): Promise<Client> {
  if (!client) {
    const connection = await Connection.connect({
      address: process.env.TEMPORAL_ADDRESS,
    });

    client = new Client({
      connection,
    });
  }

  return client;
}
