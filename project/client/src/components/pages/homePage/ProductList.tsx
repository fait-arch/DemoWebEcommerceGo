import React, { useEffect, useState } from "react";

interface Product {
	id: string;
	name: string;
	color: string;
	price: number;
	imageSrc: string;
	imageAlt: string;
}

const ProductList: React.FC = () => {
	const [products, setProducts] = useState<Product[]>([]);
	const [selectedProductIds, setSelectedProductIds] = useState<string[]>([]);

	useEffect(() => {
		fetch("http://localhost:3000/product")
			.then((response) => response.json())
			.then((data) => {
				setProducts(data);
			})
			.catch((error) => {
				console.error("Error fetching products:", error);
			});
	}, []);

	const handleAddToCart = (productId: string) => {
		if (!selectedProductIds.includes(productId)) {
			setSelectedProductIds([...selectedProductIds, productId]);
			console.log("ID del producto:", productId);
		} else {
			console.log("El producto ya está en el carrito.");
		}
	};

	const printSelectedProductIds = () => {
		const idsString = selectedProductIds.join(", ");
		console.log(`IDs de productos seleccionados: ${idsString}`);
	};

	const sendIdsToServer = () => {
		const idsString = selectedProductIds.join(",");
		fetch("http://localhost:3000/addToCart", {
			method: "POST",
			headers: {
				"Content-Type": "text/plain",
			},
			body: idsString,
		})
			.then((response) => response.text())
			.then((data) => {
				console.log(data);
				// Manejar la respuesta del servidor si es necesario
			})
			.catch((error) => {
				console.error("Error al enviar IDs de productos:", error);
			});
	};

	return (
		<div className="bg-white">
			<div className="mx-auto max-w-2xl px-4 py-16 sm:px-6 sm:py-24 lg:max-w-7xl lg:px-8">
				<h2 className="text-2xl font-bold tracking-tight text-gray-900">
					Productos
				</h2>

				<div className="mt-6 grid grid-cols-1 gap-x-6 gap-y-10 sm:grid-cols-2 lg:grid-cols-4 xl:gap-x-8">
					{products.map((product: Product) => (
						<div key={product.id} className="group relative">
							<div className="aspect-h-1 aspect-w-1 w-full overflow-hidden rounded-md bg-gray-200 lg:aspect-none group-hover:opacity-75 lg:h-80">
								<img
									src={product.imageSrc}
									alt={product.imageAlt}
									className="h-full w-full object-cover object-center lg:h-full lg:w-full"
								/>
							</div>
							<div className="mt-4">
								<div className="flex justify-between">
									<div>
										<h3 className="text-sm text-gray-700">
											{product.name}
										</h3>
										<p className="mt-1 text-sm text-gray-500">
											{product.color}
										</p>
									</div>
									<p className="text-sm font-medium text-gray-900">
										${product.price}
									</p>
								</div>
								<button
									onClick={() => {
										handleAddToCart(product.id);
										sendIdsToServer(); // Aquí es donde se llama la función correctamente
									}}
									className="mt-2 px-3 py-1 border border-transparent text-xs font-medium rounded-full shadow-sm text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
								>
									Agregar al carrito
								</button>
							</div>
						</div>
					))}
				</div>
			</div>
			<button onClick={printSelectedProductIds}>
				Imprimir IDs de productos seleccionados
			</button>
		</div>
	);
};

export default ProductList;
export const getSelectedProductIds = (selectedProductIds: string[]): string => {
	return selectedProductIds.join(",");
};
