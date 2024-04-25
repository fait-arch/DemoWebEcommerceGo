import React from "react";

import MenuBar from "../homePage/menuBar.tsx";
import Login from "./login.tsx";

export default function App(): JSX.Element {
	return (
		<div>
			<MenuBar />
			<Login />
		</div>
	);
}
