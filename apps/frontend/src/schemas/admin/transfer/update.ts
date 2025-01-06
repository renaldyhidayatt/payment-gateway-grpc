import { z } from "zod";

export const updateTransferRequestSchema = z.object({
  transfer_from: z.string().min(1).max(16),
  transfer_to: z.string().min(1).max(16),
  transfer_amount: z.number().int().min(1).max(16),
});

export type UpdateTransferFormValues = z.infer<
  typeof updateTransferRequestSchema
>;
