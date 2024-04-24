import React, { useState } from "react";
import Button from "./button";

const Counter: React.FC = () => {
	const [ids, setIds] = useState<string[]>([]);

	const handleClick = (id: string) => {
		if (!ids.includes(id)) {
			setIds([...ids, id]);
		}
	};

	return (
		<div>
			<h2>Counter</h2>
			<p>Total unique IDs: {ids.length}</p>
			{ids.map((id, index) => (
				<p key={index}>{id}</p>
			))}
			<Button id="1" onClick={handleClick} />
			<Button id="2" onClick={handleClick} />
			<Button id="3" onClick={handleClick} />
		</div>
	);
};

export default Counter;
