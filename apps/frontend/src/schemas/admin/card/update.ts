import { z } from "zod";

export const updateCardRequestSchema = z.object({
  user_id: z.number().int().min(1),
  card_type: z.string(),
  expire_date: z.coerce.date(),
  cvv: z.string(),
  card_provider: z.string(),
});

export type UpdateCardFormValues = z.infer<typeof updateCardRequestSchema>;
