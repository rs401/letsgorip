import { Post } from "./post";

export class Thread {
  Id?: number;
  ForumId?: number;
  UserId?: number;
  Title?: string;
  Msg?: string;
  Posts?: Array<Post>;
  CreatedAt?: Date;
  UpdatedAt?: Date;

  constructor() {}
}
