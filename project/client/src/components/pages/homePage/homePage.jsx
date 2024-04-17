import Banner from "./banner.jsx";
import MenuBar from "./menuBar.jsx";
import ProductListing from "./productListing.jsx";

export default function App() {
	return (
		<div>
			<MenuBar />
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

			<ProductListing />
		</div>
	);
}
