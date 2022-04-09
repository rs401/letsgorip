export class User {
  public id?: number;
  public uid?: string;
  public name?: string;
  public email?: string;
  public email_verified?: boolean;
  public picture?: string;
  public role?: number;
  public createdAt?: Date;
  public updatedAt?: Date;

  constructor() {
    
  }
  // constructor(uid: string, name: string, email: string, email_verified: boolean, picture: string = '') {
  //   this
  // }
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