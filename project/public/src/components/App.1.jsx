import React from "react";
import { BrowserRouter } from "react-router-dom";
import MenuBar from "./pages/homePage/menuBar";
import ProductListingPage from "./pages/homePage/productListingPage";
import CartComponent from "./pages/detailsProduct/cartItems.jsx";
import Banner from "./pages/homePage/banner.jsx";

// Definición de la lista de publicaciones

export default function App() {
	return (
		<div>
			<MenuBar />

			<BrowserRouter>
				<Routes></Routes>
			</BrowserRouter>

			<CartComponent />

			<Banner />

			<button
				onClick={async () => {
					const response = await fetch("http://localhost:3000/users");
					const data = await response.json();
					console.log(data);
				}}
			>
				Obtener ddd datos
			</button>

			<ProductListingPage />
		</div>
	);
}
