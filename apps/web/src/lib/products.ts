export const products: IProduct[] = [
  {
    id: 1,
    name: 'Shëlff',
    price: 12.99,
  },
  {
    id: 2,
    name: 'Kåffë Tåblë',
    price: 49.99,
  },
  {
    id: 3,
    name: 'Böökëkäsë',
    price: 79.99,
  },
  {
    id: 4,
    name: 'Kåblë Tïdy',
    price: 2.99,
  },
  {
    id: 5,
    name: 'Kändlën',
    price: 3.49,
  },
];

export function getProduct(t: number): IProduct | undefined {
  return products.find(({ id }) => id === t);
}
