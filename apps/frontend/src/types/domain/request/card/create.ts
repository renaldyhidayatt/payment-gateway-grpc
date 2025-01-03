export interface CreateCard {
    UserID: number;
    CardType: string;
    ExpireDate: Date;
    CVV: string;
    CardProvider: string;
}