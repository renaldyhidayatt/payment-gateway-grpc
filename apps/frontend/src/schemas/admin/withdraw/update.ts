import { z } from "zod";

export const updateWithdrawRequestSchema = z.object({
  card_number: z.string().min(1).max(16),
  withdraw_amount: z.number().int().min(1).max(16),
  withdraw_time: z.date(),
});

export type UpdateWithdrawFormValues = z.infer<
  typeof updateWithdrawRequestSchema
>;
