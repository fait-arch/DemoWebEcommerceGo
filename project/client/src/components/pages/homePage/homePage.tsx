import React, { useState } from "react";
import Banner from "./banner.tsx";
import MenuBar from "./menuBar.tsx";
import ProductList from "./productListing.tsx";
import SelectedProductIds from "./selectedProductIDs.tsx"; // Importa el nuevo componente

export default function App(): JSX.Element {
	// Supongamos que aquí tienes una lista de IDs de productos seleccionados
	const [selectedProductIds, setSelectedProductIds] = useState<string[]>([]);

	return (
		<div>
			<MenuBar />
			<Banner />
			<ProductList />
			{/* Renderiza el componente SelectedProductIds y pasa la lista de IDs seleccionados como prop */}
			<SelectedProductIds selectedIds={selectedProductIds} />
		</div>
	);
}
