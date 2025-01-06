import { z } from "zod";

export const updateTopupRequestSchema = z.object({
  card_number: z.string().min(1).max(16),
  topup_amount: z.number().int().min(1).max(16),
  topup_method: z.string().min(1).max(16),
});

export type UpdateTopupFormValues = z.infer<typeof updateTopupRequestSchema>;
