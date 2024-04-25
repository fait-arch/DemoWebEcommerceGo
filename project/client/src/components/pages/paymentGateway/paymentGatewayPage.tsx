import React from "react";
import MenuBar from "../homePage/menuBar";
import CartItems from "./cartItems";

const App: React.FC = (): JSX.Element => {
	return (
		<div>
			<MenuBar />
			<CartItems />
		</div>
	);
};

export default App;
