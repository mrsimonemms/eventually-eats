export const products: IProduct[] = [
  {
    id: 1,
    name: 'Chips',
    price: 3.5,
  },
  {
    id: 2,
    name: 'Battered cod',
    price: 8.75,
  },
  {
    id: 3,
    name: 'Battered haddock',
    price: 9.75,
  },
  {
    id: 4,
    name: 'Curry sauce',
    price: 1.45,
  },
  {
    id: 5,
    name: 'Gravy',
    price: 1.45,
  },
];

export function getProduct(t: number): IProduct | undefined {
  return products.find(({ id }) => id === t);
}
