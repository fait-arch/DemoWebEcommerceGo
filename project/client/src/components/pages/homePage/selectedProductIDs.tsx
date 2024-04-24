import React from "react";
import { getIdsString } from "./productListing";

const SelectedProductIds: React.FC<{ selectedIds: string[] }> = ({
	selectedIds,
}) => {
	const idsString = getIdsString(selectedIds);

	return (
		<div>
			<h2>IDs de productos seleccionados:</h2>
			<p>{idsString}</p>
		</div>
	);
};

export default SelectedProductIds;
