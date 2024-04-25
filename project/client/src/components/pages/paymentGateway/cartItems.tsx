import React, { useState, useEffect } from "react";

interface Product {
	id: number;
	name: string;
	description: string;
	price: number;
	image: string;
	quantity: number;
}

function Cart() {
	const [cartProducts, setCartProducts] = useState<Product[]>([]);
	const [subtotal, setSubtotal] = useState<number>(0);
	const [total, setTotal] = useState<number>(0);

	useEffect(() => {
		fetch("http://127.0.0.1:3000/productID")
			.then((response) => response.json())
			.then((data) => {
				// Mapear los datos de la API al formato de productos
				const productsFromAPI: Product[] = data.map((item: any) => ({
					id: item.id,
					name: item.name,
					description: item.description,
					price: item.price,
					image: item.imageSrc,
					quantity: item.quantity,
				}));
				setCartProducts(productsFromAPI);

				// Calcular el subtotal y el total
				const newSubtotal = productsFromAPI.reduce(
					(acc, curr) => acc + curr.price * curr.quantity,
					0
				);
				setSubtotal(newSubtotal);
				setTotal(newSubtotal);
			})
			.catch((error) => console.error("Error fetching products:", error));
	}, []);

	const handleRemoveItem = (productId: number, price: number) => {
		setCartProducts(
			cartProducts.filter((product) => product.id !== productId)
		);
		setSubtotal(subtotal - price);
		setTotal(total - price);
	};

	const handleQuantityChange = (productId: number, newQuantity: number) => {
		const updatedProducts = cartProducts.map((product) => {
			if (product.id === productId) {
				return {
					...product,
					quantity: newQuantity,
				};
			}
			return product;
		});

		setCartProducts(updatedProducts);

		const newSubtotal = updatedProducts.reduce(
			(acc, curr) => acc + curr.price * curr.quantity,
			0
		);
		setSubtotal(newSubtotal);
		setTotal(newSubtotal);
	};

	return (
		<section className="h-screen bg-gray-100 py-12 sm:py-16 lg:py-7">
			<div className="mx-auto px-4 sm:px-6 lg:px-8">
				<div className="flex items-center justify-center">
					<h1 className="text-2xl font-semibold text-gray-900">
						Carrito de Compra
					</h1>
				</div>

				<div className="mx-auto mt-8 max-w-md md:mt-12">
					<div className="rounded-3xl bg-white shadow-lg">
						<div className="px-4 py-6 sm:px-8 sm:py-10">
							<div className="flow-root">
								<ul className="-my-8">
									{/* Renderizamos los detalles de los productos dinámicamente */}
									{cartProducts.map((product) => (
										<li
											key={product.id}
											className="flex flex-col space-y-3 py-6 text-left sm:flex-row sm:space-x-5 sm:space-y-0"
										>
											<div className="shrink-0 relative">
												<span className="absolute top-1 left-1 flex h-6 w-6 items-center justify-center rounded-full border bg-white text-sm font-medium text-gray-500 shadow sm:-top-2 sm:-right-2">
													{product.quantity}
												</span>
												<img
													className="h-24 w-24 max-w-full rounded-lg object-cover"
													src={product.image}
													alt=""
												/>
											</div>

											<div className="relative flex flex-1 flex-col justify-between">
												<div className="sm:col-gap-5 sm:grid sm:grid-cols-2">
													<div className="pr-8 sm:pr-5">
														<p className="text-base font-semibold text-gray-900">
															{product.name}
														</p>
														<p className="mx-0 mt-1 mb-0 text-sm text-gray-400">
															{
																product.description
															}
														</p>
														<div className="flex items-center">
															<label
																htmlFor={`quantity_${product.id}`}
																className="mr-2"
															></label>
															<input
																type="number"
																id={`quantity_${product.id}`}
																value={
																	product.quantity
																}
																min={1}
																onChange={(e) =>
																	handleQuantityChange(
																		product.id,
																		parseInt(
																			e
																				.target
																				.value
																		)
																	)
																}
																className="border border-gray-300 rounded-md px-3 py-1 w-16 focus:outline-none focus:border-gray-500"
															/>
														</div>
													</div>

													<div className="mt-4 flex items-end justify-between sm:mt-0 sm:items-start sm:justify-end">
														<p className="shrink-0 w-20 text-base font-semibold text-gray-900 sm:order-2 sm:ml-8 sm:text-right">
															$
															{product.price.toFixed(
																2
															)}
														</p>
													</div>
												</div>

												<div className="absolute top-0 right-0 flex sm:bottom-0 sm:top-auto">
													<button
														type="button"
														className="flex rounded p-2 text-center text-gray-500 transition-all duration-200 ease-in-out focus:shadow hover:text-gray-900"
														onClick={() =>
															handleRemoveItem(
																product.id,
																product.price *
																	product.quantity
															)
														}
													>
														<svg
															className="h-5 w-5"
															xmlns="http://www.w3.org/2000/svg"
															fill="none"
															viewBox="0 0 24 24"
															stroke="currentColor"
														>
															<path
																strokeLinecap="round"
																strokeLinejoin="round"
																strokeWidth="2"
																d="M6 18L18 6M6 6l12 12"
															/>
														</svg>
													</button>
												</div>
											</div>
										</li>
									))}
								</ul>
							</div>

							{/* Detalles del subtotal y envío */}
							<div className="mt-6 space-y-3 border-t border-b py-8">
								<div className="flex items-center justify-between">
									<p className="text-gray-400">Subtotal</p>
									<p className="text-lg font-semibold text-gray-900">
										${subtotal.toFixed(2)}
									</p>
								</div>
							</div>

							{/* Total */}
							<div className="mt-6 flex items-center justify-between">
								<p className="text-sm font-medium text-gray-900">
									Total
								</p>
								<p className="text-2xl font-semibold text-gray-900">
									<span className="text-xs font-normal text-gray-400">
										USD
									</span>{" "}
									{total.toFixed(2)}
								</p>
							</div>

							{/* Botón de realizar pedido */}
							<div className="mt-6 text-center">
								<button
									type="button"
									className="group inline-flex w-full items-center justify-center rounded-md bg-gray-900 px-6 py-4 text-lg font-semibold text-white transition-all duration-200 ease-in-out focus:shadow hover:bg-gray-200 hover:text-gray-900"
								>
									Pagar
									<svg
										xmlns="http://www.w3.org/2000/svg"
										className="group-hover:ml-8 ml-4 h-6 w-6 transition-all"
										fill="none"
										viewBox="0 0 24 24"
										stroke="currentColor"
										strokeWidth="2"
									>
										<path
											strokeLinecap="round"
											strokeLinejoin="round"
											d="M13 7l5 5m0 0l-5 5m5-5H6"
										/>
									</svg>
								</button>
							</div>
						</div>
					</div>
				</div>
			</div>
		</section>
	);
}

export default Cart;
