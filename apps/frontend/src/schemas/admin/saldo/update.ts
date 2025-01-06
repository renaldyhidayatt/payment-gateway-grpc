import { z } from "zod";

export const updateSaldoRequestSchema = z.object({
  card_number: z.string().min(1).max(16),
  total_balance: z.number().int().min(1).max(16),
});

export type UpdateSaldoFormValues = z.infer<typeof updateSaldoRequestSchema>;
