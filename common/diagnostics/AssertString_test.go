package diagnostics_test

import (
	"github.com/microbusinesses/Micro-Businesses-Core/common/diagnostics"
	"github.com/microbusinesses/Micro-Businesses-Core/system"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

var _ = Describe("Assert String is empty", func() {
	var (
		emptyString   string
		testString    string
		testValueName string
		testMessage   string
	)

	BeforeEach(func() {
		emptyString = ""

		var uuid system.UUID

		uuid, _ = system.RandomUUID()
		testString = uuid.String()

		uuid, _ = system.RandomUUID()
		testValueName = uuid.String()

		uuid, _ = system.RandomUUID()
		testMessage = uuid.String()
	})

	Describe("when empty string value provided", func() {
		It("Should return empty string", func() {
			Expect(diagnostics.StringIsEmpty(emptyString, emptyString, emptyString)).To(Equal(emptyString))
		})
	})

	Describe("when string value provided", func() {
		Context("with value name and message both emptyd", func() {
			It("Should panic with empty string as message", func() {
				defer func() {
					if r := recover(); r == nil {
						Fail("Should have paniced with empty as message!!!")
					} else if r != emptyString {
						Fail("Should have paniced with empty as message!!!")
					}
				}()

				diagnostics.StringIsEmpty(testString, emptyString, emptyString)
			})
		})

		Context("with value name provided", func() {
			It("Should panic with provided value name as message", func() {
				defer func() {
					if r := recover(); r == nil {
						Fail("Should have paniced with value name as message!!!")
					} else if r != testValueName {
						Fail("Should have paniced with value name as message!!!")
					}
				}()

				diagnostics.StringIsEmpty(testString, testValueName, emptyString)
			})
		})

		Context("with message provided", func() {
			It("Should panic with provided message as message", func() {
				defer func() {
					if r := recover(); r == nil {
						Fail("Should have paniced with message as message!!!")
					} else if r != testMessage {
						Fail("Should have paniced with message as message!!!")
					}
				}()

				diagnostics.StringIsEmpty(testString, emptyString, testMessage)
			})
		})

		Context("with name and message both provided", func() {
			It("Should panic with provided message as message", func() {
				defer func() {
					if r := recover(); r == nil {
						Fail("Should have paniced with message as message!!!")
					} else if r != testMessage {
						Fail("Should have paniced with message as message!!!")
					}
				}()

				diagnostics.StringIsEmpty(testString, testValueName, testMessage)
			})
		})
	})
})

var _ = Describe("Assert String is not empty", func() {
	var (
		emptyString   string
		testString    string
		testValueName string
		testMessage   string
	)

	BeforeEach(func() {
		emptyString = ""

		var uuid system.UUID

		uuid, _ = system.RandomUUID()
		testString = uuid.String()

		uuid, _ = system.RandomUUID()
		testValueName = uuid.String()

		uuid, _ = system.RandomUUID()
		testMessage = uuid.String()
	})

	Describe("when string value provided", func() {
		It("Should return provided string as value", func() {
			Expect(diagnostics.StringIsNotEmpty(testString, emptyString, emptyString)).To(Equal(testString))
		})
	})

	Describe("when empty string value provided", func() {
		Context("with value name and message both emptyd", func() {
			It("Should panic with empty string as message", func() {
				defer func() {
					if r := recover(); r == nil {
						Fail("Should have paniced with empty as message!!!")
					} else if r != emptyString {
						Fail("Should have paniced with empty as message!!!")
					}
				}()

				diagnostics.StringIsNotEmpty(emptyString, emptyString, emptyString)
			})
		})

		Context("with value name provided", func() {
			It("Should panic with provided value name as message", func() {
				defer func() {
					if r := recover(); r == nil {
						Fail("Should have paniced with value name as message!!!")
					} else if r != testValueName {
						Fail("Should have paniced with value name as message!!!")
					}
				}()

				diagnostics.StringIsNotEmpty(emptyString, testValueName, emptyString)
			})
		})

		Context("with message provided", func() {
			It("Should panic with provided message as message", func() {
				defer func() {
					if r := recover(); r == nil {
						Fail("Should have paniced with message as message!!!")
					} else if r != testMessage {
						Fail("Should have paniced with message as message!!!")
					}
				}()

				diagnostics.StringIsNotEmpty(emptyString, emptyString, testMessage)
			})
		})

		Context("with name and message both provided", func() {
			It("Should panic with provided message as message", func() {
				defer func() {
					if r := recover(); r == nil {
						Fail("Should have paniced with message as message!!!")
					} else if r != testMessage {
						Fail("Should have paniced with message as message!!!")
					}
				}()

				diagnostics.StringIsNotEmpty(emptyString, testValueName, testMessage)
			})
		})
	})
})

var _ = Describe("Assert String is not empty or whitespace", func() {
	var (
		emptyString   string
		testString    string
		whitespace    string
		testValueName string
		testMessage   string
	)

	BeforeEach(func() {
		emptyString = ""
		whitespace = "    "

		var uuid system.UUID

		uuid, _ = system.RandomUUID()
		testString = uuid.String()

		uuid, _ = system.RandomUUID()
		testValueName = uuid.String()

		uuid, _ = system.RandomUUID()
		testMessage = uuid.String()
	})

	Describe("when string value provided", func() {
		It("Should return provided string as value", func() {
			Expect(diagnostics.StringIsNotEmptyOrWhitespace(testString, emptyString, emptyString)).To(Equal(testString))
		})
	})

	Describe("when empty string value provided", func() {
		Context("with value name and message both emptyd", func() {
			It("Should panic with empty string as message", func() {
				defer func() {
					if r := recover(); r == nil {
						Fail("Should have paniced with empty as message!!!")
					} else if r != emptyString {
						Fail("Should have paniced with empty as message!!!")
					}
				}()

				diagnostics.StringIsNotEmptyOrWhitespace(emptyString, emptyString, emptyString)
			})
		})

		Context("with value name provided", func() {
			It("Should panic with provided value name as message", func() {
				defer func() {
					if r := recover(); r == nil {
						Fail("Should have paniced with value name as message!!!")
					} else if r != testValueName {
						Fail("Should have paniced with value name as message!!!")
					}
				}()

				diagnostics.StringIsNotEmptyOrWhitespace(emptyString, testValueName, emptyString)
			})
		})

		Context("with message provided", func() {
			It("Should panic with provided message as message", func() {
				defer func() {
					if r := recover(); r == nil {
						Fail("Should have paniced with message as message!!!")
					} else if r != testMessage {
						Fail("Should have paniced with message as message!!!")
					}
				}()

				diagnostics.StringIsNotEmptyOrWhitespace(emptyString, emptyString, testMessage)
			})
		})

		Context("with name and message both provided", func() {
			It("Should panic with provided message as message", func() {
				defer func() {
					if r := recover(); r == nil {
						Fail("Should have paniced with message as message!!!")
					} else if r != testMessage {
						Fail("Should have paniced with message as message!!!")
					}
				}()

				diagnostics.StringIsNotEmptyOrWhitespace(emptyString, testValueName, testMessage)
			})
		})
	})

	Describe("when string contains whitespace characters only provided", func() {
		Context("with value name and message both emptyd", func() {
			It("Should panic with empty string as message", func() {
				defer func() {
					if r := recover(); r == nil {
						Fail("Should have paniced with empty as message!!!")
					} else if r != emptyString {
						Fail("Should have paniced with empty as message!!!")
					}
				}()

				diagnostics.StringIsNotEmptyOrWhitespace(whitespace, emptyString, emptyString)
			})
		})

		Context("with value name provided", func() {
			It("Should panic with provided value name as message", func() {
				defer func() {
					if r := recover(); r == nil {
						Fail("Should have paniced with value name as message!!!")
					} else if r != testValueName {
						Fail("Should have paniced with value name as message!!!")
					}
				}()

				diagnostics.StringIsNotEmptyOrWhitespace(whitespace, testValueName, emptyString)
			})
		})

		Context("with message provided", func() {
			It("Should panic with provided message as message", func() {
				defer func() {
					if r := recover(); r == nil {
						Fail("Should have paniced with message as message!!!")
					} else if r != testMessage {
						Fail("Should have paniced with message as message!!!")
					}
				}()

				diagnostics.StringIsNotEmptyOrWhitespace(whitespace, emptyString, testMessage)
			})
		})

		Context("with name and message both provided", func() {
			It("Should panic with provided message as message", func() {
				defer func() {
					if r := recover(); r == nil {
						Fail("Should have paniced with message as message!!!")
					} else if r != testMessage {
						Fail("Should have paniced with message as message!!!")
					}
				}()

				diagnostics.StringIsNotEmptyOrWhitespace(whitespace, testValueName, testMessage)
			})
		})
	})
})

func TestAssertString(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Assert String is empty")
	RunSpecs(t, "Assert String is not empty")
	RunSpecs(t, "Assert String is not empty or whitespace")
}