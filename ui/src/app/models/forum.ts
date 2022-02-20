import { Thread } from "./thread";

export class Forum {
  Id?: number;
  UserId?: number;
  Title?: string;
  Description?: string;
  Threads?: Array<Thread>;
  CreatedAt?: Date;
  UpdatedAt?: Date;

  constructor() {}
}
