import React from "react";

interface ButtonProps {
	id: string;
	onClick: (id: string) => void;
}

const Button: React.FC<ButtonProps> = ({ id, onClick }) => {
	const handleClick = () => {
		onClick(id);
	};

	return <button onClick={handleClick}>Click me</button>;
};

export default Button;
