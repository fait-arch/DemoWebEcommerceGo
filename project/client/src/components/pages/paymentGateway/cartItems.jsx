import React, { useState } from "react";

// Definimos un objeto JSON que contiene los detalles del producto
const products = [
	{
		id: 1,
		name: "Nike Air Max 2019",
		description: "36EU - 4US",
		price: 1259.0,
		image: "https://images.unsplash.com/photo-1588484628369-dd7a85bfdc38?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxzZWFyY2h8MTh8fHNuZWFrZXJ8ZW58MHx8MHx8&auto=format&fit=crop&w=150&q=60",
	},
	{
		id: 2,
		name: "Product 2",
		description: "Description of Product 2",
		price: 1000.0,
		image: "https://i.pinimg.com/564x/71/5e/94/715e943539b1caa889f39bfc1a36f4aa.jpg",
	},
	// Add more products as needed
];

function Cart() {
	const [cartProducts, setCartProducts] = useState(products);
	const [subtotal, setSubtotal] = useState(
		products.reduce((acc, curr) => acc + curr.price, 0)
	);
	const [total, setTotal] = useState(subtotal);

	const handleRemoveItem = (productId, price) => {
		setCartProducts(
			cartProducts.filter((product) => product.id !== productId)
		);
		setSubtotal(subtotal - price);
		setTotal(total - price);
	};

	return (
		<section className="h-screen bg-gray-100 py-12 sm:py-16 lg:py-20">
			<div className="mx-auto px-4 sm:px-6 lg:px-8">
				<div className="flex items-center justify-center">
					<h1 className="text-2xl font-semibold text-gray-900">
						Your Cart
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
													1
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
																product.price
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
									className="group inline-flex w-full items-center justify-center rounded-md bg-orange-500 px-6 py-4 text-lg font-semibold text-white transition-all duration-200 ease-in-out focus:shadow hover:bg-gray-800"
								>
									Place Order
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
