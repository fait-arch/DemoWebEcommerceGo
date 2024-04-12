import React from "react";

import MenuBar from "./MenuBar";

// Definición de la lista de publicaciones

export default function App() {
	return (
		<div>
			<MenuBar />
			<button
				onClick={async () => {
					const response = await fetch("http://localhost:3000/users");
					const data = await response.json();
					console.log(data);
				}}
			>
				Obtener ddd datos
			</button>
		</div>
	);
}
