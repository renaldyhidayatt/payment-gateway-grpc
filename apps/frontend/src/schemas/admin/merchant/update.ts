import { z } from "zod";

export const updateMerchantRequestSchema = z.object({
  name: z.string(),
  user_id: z.number().int().min(1),
  status: z.string(),
});

export type UpdateMerchantFormValues = z.infer<
  typeof updateMerchantRequestSchema
>;
