import React from "react";
import MenuBar from "../homePage/menuBar";
import FilterProductCollection from "./filterProducts";

export default function App(): JSX.Element {
	return (
		<div>
			<MenuBar />
			<FilterProductCollection />
		</div>
	);
}
