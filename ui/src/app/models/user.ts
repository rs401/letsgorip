export class User {
  public ID?: number;
  public Name?: string;
  public Email?: string;
  public Role?: number;
  public CreatedAt?: Date;
  public UpdatedAt?: Date;

  constructor() {}
}

export class SignIn {
  public Email?: string;
  public Password?: string;

  constructor() {}
}