import React from "react";
import { BrowserRouter, Route, Routes } from "react-router-dom";

import HomePage from "./pages/homePage/homePage";
import Login from "./pages/loginPage/loginPage";
import PaymentGateway from "./pages/paymentGateway/paymentGatewayPage";
import AllProductsPage from "./pages/allProducts/allProductsPage";

// Definición de la lista de publicaciones

export default function App() {
	return (
		<BrowserRouter>
			<Routes>
				<Route path="/" element={<HomePage />} />
				<Route path="/paymentGateway" element={<PaymentGateway />} />
				<Route path="/login" element={<Login />} />
				<Route path="/allProductsPage" element={<AllProductsPage />} />
			</Routes>
		</BrowserRouter>
	);
}
