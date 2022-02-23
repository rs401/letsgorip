import { Thread } from "./thread";

export class Forum {
  id?: number;
  user_id?: number;
  title?: string;
  description?: string;
  threads?: Array<Thread>;
  created_at?: Date;
  updated_at?: Date;

  constructor() {}
}
