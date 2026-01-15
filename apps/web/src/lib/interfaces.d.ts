interface IProduct {
  id: number;
  name: string;
  price: number;
  quantity?: number;
}

interface IProduct2 {
  collection: boolean;
  products: { productId: number; quantity: number }[];
  status: OrderStatus;
}

type OrderStatus =
  | 'DEFAULT' // Order not paid yet
  | 'PENDING' // Order paid and waiting for restaurant to accept
  | 'ACCEPTED' // Restaurant accepted order, but not started work yet
  | 'PREPARING' // Restaurant is cooking your food
  | 'READY' // Food is ready for collection/out for delivery
  | 'REJECTED' // Kitchen has rejected the order
  | 'COMPLETED'; // Food given to a hungry person

interface IOrderState {
  collection: boolean;
  products: IProduct[];
  status: OrderStatus;
}

interface IOrder {
  orderId: string;
  state: IOrderState;
  created: Date;
}
