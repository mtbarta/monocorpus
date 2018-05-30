import validator from 'validator'
import owasp from 'owasp-password-strength-test'
import UserData from './userData'

owasp.config({
  allowPassphrases       : false,
  maxLength              : 128,
  minLength              : 10,
  minPhraseLength        : 20,
  minOptionalTestsToPass : 4,
});

interface Validator {
  (input: string): ValidationResults
}

/**
 * outputs a string if there's an error. otherwise null.
 * @param email 
 */
export const EmailValidator = (email: string): ValidationResults => {
  const isEmail = validator.isEmail(email)
  const errs = isEmail ? [] : ['Input is not a valid email.']

  return {
    errors: errs,
    failedTests: [],
    passedTests: [],
    requiredTestErrors: [],
    optionalTestErrors: [],
    isPassphrase: false,
    strong: false,
    optionalTestsPassed: 0
  }
}

export interface ValidationResults {
  errors: string[]
  failedTests: number[]
  passedTests: number[]

  requiredTestErrors: string[]
  optionalTestErrors: string[]

  isPassphrase: boolean
  strong: boolean
  optionalTestsPassed: number
}

export const PasswordValidator = (pw: string) : ValidationResults => {
  return owasp.test(pw)
}

export const PasswordExplanations = [
  'Must be 10 characters long',
  'Must be less than 128 characters',
  'Must not repeat 3 characters in a row',
  'Must contain at least one lowercase letter',
  'Must contain at least one uppercase letter',
  'Must contain at least one number',
  'Must contain at least one special character'
]

// export default class UserDataValidator {
//   emailValidator: Validator
//   passwordValidtor: Validator

//   constructor() {
//     this.emailValidator = EmailValidator
//     this.passwordValidtor = PasswordValidator
//   }

//   validate(userData: UserData):
// }