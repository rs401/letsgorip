export class User {
  public id?: number;
  public name?: string;
  public email?: string;
  public role?: number;
  public createdAt?: Date;
  public updatedAt?: Date;

  constructor() {}
}

export class SignIn {
  public email?: string;
  public password?: string;

  constructor(email: string, password: string) {
    this.email = email;
    this.password = password;
  }
}
export class SignUp {
  public name?: string;
  public email?: string;
  public password?: string;

  constructor(name: string, email: string, password: string) {
    this.name = name;
    this.email = email;
    this.password = password;
  }
}