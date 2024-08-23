export interface LoginCredentials {
	email: string;
	password: string;
}

export interface RegisterData extends LoginCredentials {
	confirmPassword: string;
}

export interface MenuItem {
	label: string;
	icon: JSX.Element;
	href?: string;
	onClick?: () => void;
}

export type Listing = {
	id: string;
	title: string;
	description: string;
	price: number;
	imageUrl: string;
};

export interface ListingState {
	listings: Listing[];
	addListing: (listing: Listing) => void;
	removeListing: (id: string) => void;
	updateListing: (id: string, updatedListing: Partial<Listing>) => void;
}

export interface AuthContextType {
	isAuthenticated: boolean;
	login: (credentials: LoginCredentials) => Promise<void>;
	logout: () => void;
}
