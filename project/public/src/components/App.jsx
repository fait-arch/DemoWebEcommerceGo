import React from "react";
import { BrowserRouter, Route, Routes } from "react-router-dom";

import MenuBar from "./pages/homePage/menuBar";
import ProductListingPage from "./pages/homePage/productListingPage";
import Demop from "./pages/loginPage/login";
import CartComponent from "./pages/detailsProduct/cartItems.jsx";
import Banner from "./pages/homePage/banner.jsx";

// Definición de la lista de publicaciones

export default function App() {
	return (
		<BrowserRouter>
			<Routes>
				<Route path="" element={<MenuBar />} />
				<Route path="/banner" element={<Banner />} />
			</Routes>
		</BrowserRouter>
	);
}
