import React, { useState } from "react";
import Banner from "./banner.tsx";
import MenuBar from "./menuBar.tsx";
import ProductListing from "./ProductList.tsx";

export default function App(): JSX.Element {
	// Supongamos que aquí tienes una lista de IDs de productos seleccionados
	const [selectedProductIds, setSelectedProductIds] = useState<string[]>([]);

	return (
		<div>
			<MenuBar />
			<Banner />
			<ProductListing />
		</div>
	);
}
