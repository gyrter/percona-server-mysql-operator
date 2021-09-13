package cluster

import (
	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	corev1 "k8s.io/api/core/v1"

	v2 "github.com/percona/percona-mysql/api/v2"
	"github.com/percona/percona-mysql/pkg/database/mysql"
	"github.com/percona/percona-mysql/pkg/k8s"
)

func (r *MySQLReconciler) reconcileMySQL(log logr.Logger, cr *v2.PerconaServerForMySQL) error {
	m := mysql.New(cr)
	sfs := m.StatefulSet()

	initImage, err := k8s.InitImage(r.Client, cr)
	if err != nil {
		return errors.Wrap(err, "get init image")
	}
	sfs.Spec.Template.Spec.InitContainers = []corev1.Container{m.InitContainer(initImage)}

	if err := k8s.SetControllerReference(cr, sfs, r.Scheme); err != nil {
		return errors.Wrap(err, "get init image")
	}

	if err := r.createOrUpdate(sfs); err != nil {
		return errors.Wrap(err, "create or update mysql statefulset")
	}

	return nil
}
