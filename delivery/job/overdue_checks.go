package job

import (
	"fmt"

	"git.teqnological.asia/teq-go/teq-echo/client/mysql"
	"git.teqnological.asia/teq-go/teq-pkg/teqlogger"
	"git.teqnological.asia/teq-go/teq-pkg/teqsentry"
	"github.com/pkg/errors"
	"github.com/robfig/cron/v3"
	// "github.com/robfig/cron/v3"
)

type borrowingOverdueCheck struct {
	// db *sql.DB
}

func NewBorrowingOverdueCheck() IJob {
	return borrowingOverdueCheck{
		// db: db,
	}
}

func (b borrowingOverdueCheck) callBorrowingOverdueCheck() {
	var users_id []int64
	if err := mysql.GetDB().Table("borrows").Select("user_id").Where("DATEDIFF(NOW(), borrow_date) > 14").Find(&users_id).Error; err != nil {
		fmt.Println("Error while getting borrow_date")
	}

	fmt.Println("Users is overdue borrowing book:", users_id)
}

func (b borrowingOverdueCheck) Run() {
	c := cron.New()

	// _, err := c.AddFunc("0 8 * * *", b.callBorrowingOverdueCheck)
	_, err := c.AddFunc("@every 100s", b.callBorrowingOverdueCheck)
	if err != nil {
		teqlogger.GetLogger().Fatal("failed to schedule borrowing overdue check")
		teqsentry.Fatal(errors.Wrap(err, "failed to schedule borrowing overdue check"))
	}

	c.Start()
}
