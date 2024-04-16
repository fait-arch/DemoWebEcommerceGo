import { Outlet } from "react-router-dom";

const Layout = () => {
	return (
		<div>
			<ul>
				<li>
					{" "}
					<Link to="/">Home</Link>{" "}
				</li>
				<li>
					{" "}
					<Link to="/#">Productos</Link>{" "}
				</li>
				<li>
					{" "}
					<Link to="/carrito">Carrito</Link>{" "}
				</li>
			</ul>

			<nav></nav>
			<hr />
			<Outlet />
		</div>
	);
};

export default Layout;
