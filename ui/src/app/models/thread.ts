import { Post } from "./post";

export class Thread {
  id?: number;
  forum_id?: number;
  user_id?: number;
  title?: string;
  msg?: string;
  posts?: Array<Post>;
  created_at?: Date;
  updated_at?: Date;

  constructor() {}
}
