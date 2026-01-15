import { ensureConnection } from '$lib/server/temporal';
import { json, type RequestHandler } from '@sveltejs/kit';

export const GET: RequestHandler = async ({ params }) => {
  const temporal = await ensureConnection();

  const handler = temporal.workflow.getHandle(params.orderId ?? '');

  // Validate the workflow exists
  await handler.describe();

  return json(await handler.query('GET_STATUS'));
};

export const POST: RequestHandler = async ({ params, request }) => {
  const temporal = await ensureConnection();

  const data = await request.json();
  const handler = temporal.workflow.getHandle(params.orderId ?? '');

  // Validate the workflow exists
  await handler.describe();

  await handler.executeUpdate('UPDATE_STATUS', {
    args: [data.status],
  });

  return json({
    hello: 'world',
  });
};
