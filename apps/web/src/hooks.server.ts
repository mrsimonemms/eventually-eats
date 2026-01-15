import { ensureConnection } from '$lib/server/temporal';
import type { ServerInit } from '@sveltejs/kit';
import process from 'node:process';

export const init: ServerInit = async () => {
  console.log('Connecting to Temporal service');
  const temporal = await ensureConnection();

  process.on('sveltekit:shutdown', async () => {
    console.log('Closing Temporal connection');
    await temporal.connection.close();
  });
};
