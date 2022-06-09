describe('The profile page', () => {

    it('display information of user in the input', () => {
        cy.visit('http://localhost:3000/profile') // change URL to match your dev URL
        cy.wait(500)

        //check value
        cy.fixture('userInformations').then((userInformationsForLogin) => {
            //Lastname
            cy.get('#firstname').should('have.value', userInformationsForLogin.firstname)
            //Lastname
            cy.get('#lastname').should('have.value', userInformationsForLogin.lastname)
            //email
            cy.get('#email').should('have.value', userInformationsForLogin.email)
            //mobile
            cy.get('#mobile').should('have.value', userInformationsForLogin.mobile)
        })
    })





    it('change Password', () => {
        cy.visit('http://localhost:3000/profile') // change URL to match your dev URL
        cy.wait(500)

        cy.fixture('userInformations').then((userInformationsForRegister) => {
            cy.get('#buttonChangePassword').click()
            cy.wait(500)

            //check if we have an error message with a wrong password
            cy.get('#passwordVerif').type(userInformationsForRegister.wrongPassword)
            cy.wait(1000)
            //click on the page to display the error message
            cy.get('#bodyOfThePage').click()
            cy.get('#errorMessageWrongPassword').should('have.text', 'the password is not good')
            cy.wait(1000)

            //change the password
            cy.get('#passwordVerif').clear()
            cy.get('#passwordVerif').type(userInformationsForRegister.password)
            cy.wait(2000)

            //check wrong new password
            cy.get('#newPassword').type(userInformationsForRegister.otherPassword)
            cy.wait(1000)
            cy.get('#bodyOfThePage').click()
            cy.get('#errorMessageWrongSamePassword').should('have.text', 'the password must be at least 8 characters long' )

            cy.get('#newPassword').clear()
            cy.get('#newPassword').type(userInformationsForRegister.newPassword)

            //check wrong verif password
            cy.get('#newPasswordCheck').type(userInformationsForRegister.wrongPassword)
            cy.wait(1000)
            //click on the page to display the error message
            cy.get('#bodyOfThePage').click()
            cy.get('#errorMessageWrongSamePassword').should('have.text', 'the passwords are not the same')
            cy.wait(1000)

            //check good verif password
            cy.get('#newPasswordCheck').clear()
            cy.get('#newPasswordCheck').type(userInformationsForRegister.newPassword)
            cy.get('#bodyOfThePage').click()
            cy.wait(1000)
            cy.get('#messagePasswordsAreTheSame').should('have.text', 'the Passwards are the same ! GOOD JOB')
        })
    })




    it('check disable SAVE password', () => {
        cy.visit('http://localhost:3000/profile') // change URL to match your dev URL
        cy.wait(500)

        //check the save button is disable
        cy.get('#saveChange').should('be.disabled')

        cy.fixture('userInformations').then((userInformationsForRegister) => {
            cy.get('#firstname').clear()
            cy.get('#firstname').type(userInformationsForRegister.newFirstname)
            cy.wait(1000)
            cy.get('#bodyOfThePage').click()

            //the button is not disable because we have changed a value
            cy.get('#saveChange').should('not.be.disabled')
        })
    })





    it('new firstname and new lastname new email new mobile values are good', () => {
        cy.visit('http://localhost:3000/profile') // change URL to match your dev URL
        cy.wait(500)

        cy.fixture('userInformations').then((userInformationsForRegister) => {
            //firstname too long
            cy.get('#firstname').clear()
            cy.get('#firstname').type(userInformationsForRegister.fakeName_128)
            cy.get('#bodyOfThePage').click()
            cy.get('#errorMessageLengthNewFirstName').contains('Le nouveau prénom est trop long')

            //lastname too long
            cy.get('#lastname').clear()
            cy.get('#lastname').type(userInformationsForRegister.fakeName_128)
            cy.get('#bodyOfThePage').click()
            cy.get('#errorMessageLengthNewFirstName').contains('Le nouveau nom est trop long')

            //email format is not valid
            cy.get('#email').clear()
            cy.get('#email').type(userInformationsForRegister.fakeEmail)
            cy.get('#bodyOfThePage').click()
            cy.get('#errorMessageIncorrectEmail').contains('The email is not valid')

            //mobile format is not valid
            cy.get('#mobile').clear()
            cy.get('#mobile').type(userInformationsForRegister.fakePhone)
            cy.get('#bodyOfThePage').click()
            cy.get('#errorMessageLengthNewFirstName').contains('The mobile is not valid')
        })
    })
})