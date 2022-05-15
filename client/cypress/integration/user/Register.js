describe('The Register page', () => {


    it('check creation of a account', () => {
        cy.visit('http://localhost:3000/register') // change URL to match your dev URL
        cy.wait(500)
        //enter the information in the inputs
        cy.fixture('userInformations').then((userInformationsForRegister) => {
            cy.get('#firstname').type(userInformationsForRegister.firstname)
            cy.get('#lastname').type(userInformationsForRegister.lastname)
            cy.get('#email').type(userInformationsForRegister.email)
            cy.get('#mobile').type(userInformationsForRegister.mobile)
            cy.get('#password').type(userInformationsForRegister.password)
            // test if it's the same password
            cy.get('#passwordCheck').type(userInformationsForRegister.password)

            cy.log(userInformationsForRegister.firstname)
        })
        cy.wait(500)
        //click on register button for confirm the creation of the account
        cy.get('#registerButton').click()
        cy.wait(500)
        //check if we are redirected to the home page
        cy.url().should('eq', 'http://localhost:3000')
        cy.log('à voir si sa suffit de voir si on est redirigé pour savoir si on est connecté')
    })





    it('check Register button is disable if all informations have not been entered', () => {
        cy.visit('http://localhost:3000/register')
        cy.wait(500)
        //register button is disable because the input are empty
        cy.get('#registerButton').should('be.disabled')
        cy.wait(500)
        //enter information to test if the register button is able now
        cy.fixture('userInformations').then((userInformationsForRegister) => {
            cy.get('#firstname').type(userInformationsForRegister.firstname)
            cy.wait(500)
            cy.get('#registerButton').should('be.disabled')
            cy.get('#lastname').type(userInformationsForRegister.lastname)
            cy.wait(500)
            cy.get('#registerButton').should('be.disabled')
            cy.get('#email').type(userInformationsForRegister.email)
            cy.wait(500)
            cy.get('#registerButton').should('be.disabled')
            cy.get('#mobile').type(userInformationsForRegister.mobile)
            cy.wait(500)
            cy.get('#registerButton').should('be.disabled')
            cy.get('#password').type(userInformationsForRegister.password)
            cy.wait(500)
            cy.get('#registerButton').should('be.disabled')
            // test if it's the same password
            cy.get('#passwordCheck').type(userInformationsForRegister.password)

            cy.wait(1000)
            cy.get('#registerButton').should('not.be.disabled') //all input are good

            //check if clear a input
            cy.get('#firstname').clear()
            cy.wait(1000)
            cy.get('#registerButton').should('be.disabled')
        })
    })





    it('firstname and lastname are less than 128 characters', () => {
        cy.visit('http://localhost:3000/register')
        cy.wait(500)
        cy.fixture('userInformations').then((userInformationsForRegister) => {
            cy.get('#firstname').type(userInformationsForRegister.fakeName_128)
            cy.wait(1000)
            cy.get('#errorMessageLengthFirstName').contains('Le prénom est trop long')

            cy.wait(1000)
            cy.get('#lastname').type(userInformationsForRegister.fakeName_128)
            cy.wait(1000)
            cy.get('#errorMessageLengthLastName').contains('Le nom est trop long')
        })
    })





    it('email input only takes email', () => {
        cy.visit('http://localhost:3000/register')
        cy.wait(500)
        cy.fixture('userInformations').then((userInformationsForRegister) => {
            cy.get('#email').type(userInformationsForRegister.fakeEmail)
        })
        cy.wait(1000)
        cy.get('#errorMessageIncorrectEmail').contains('The email is not valid')
    })





    it('mobile take only numbers value', () => {
        cy.visit('http://localhost:3000/register')
        cy.wait(500)
        cy.fixture('userInformations').then((userInformationsForRegister) => {
            cy.get('#mobile').type(userInformationsForRegister.fakePhone)
        })
        cy.wait(1000)
        cy.get('#errorMessageIncorrectMobile').contains('The mobile is not valid')
    })





    it('2 inputs for passward have the same value', () => {
        cy.visit('http://localhost:3000/register')
        cy.wait(500)
        cy.fixture('userInformations').then((userInformationsForRegister) => {
            cy.get('#password').type(userInformationsForRegister.password)
            cy.get('#passwordCheck').type(userInformationsForRegister.otherPassword)
            cy.wait(1000)
            cy.get('#errorMessageIncorrectMobile').contains('The password is not the same for the inputs')
        })
    })





    it('password is longer than 8 characters', () => {
        cy.visit('http://localhost:3000/register')
        cy.wait(500)
        cy.fixture('userInformations').then((userInformationsForRegister) => {
            cy.get('#password').type(userInformationsForRegister.otherPassword)
            cy.wait(1000)
            cy.get('#errorMessageIncorrectMobile').contains('the password must be at least 8 characters long')
        })
    })
})