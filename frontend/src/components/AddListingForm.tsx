import { useState } from "react";
import useListingStore from "../store/useListingStore";

const AddListingForm = () => {
	const [title, setTitle] = useState("");
	const addListing = useListingStore((state) => state.addListing);

	const handleSubmit = (event: React.FormEvent) => {
		event.preventDefault();
		addListing({
			id: Date.now().toString(),
			title,
			description: "",
			price: 0,
			imageUrl: "",
		});
		setTitle("");
	};

	return (
		<form onSubmit={handleSubmit}>
			<input
				type="text"
				value={title}
				onChange={(e) => setTitle(e.target.value)}
				placeholder="Listing title"
				required
			/>
			<button type="submit">Add Listing</button>
		</form>
	);
};

export default AddListingForm;
