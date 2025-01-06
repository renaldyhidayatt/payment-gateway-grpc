import { z } from "zod";

export const createTransactionRequestSchema = z.object({
  card_number: z.string().min(1).max(16),
  amount: z.number().min(1).max(16),
  payment_method: z.string().min(1).max(16),
  merchant_id: z.number().min(1),
  transaction_time: z.date(),
});

export type CreateTransactionFormValues = z.infer<
  typeof createTransactionRequestSchema
>;
