// Copyright 2016 The Prometheus Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package collector

import (
	"os"
	"testing"
)

func TestArcstatsParsing(t *testing.T) {
	arcstatsFile, err := os.Open("fixtures/proc/spl/kstat/zfs/arcstats")
	if err != nil {
		t.Fatal(err)
	}
	defer arcstatsFile.Close()

	c := zfsCollector{}
	if err != nil {
		t.Fatal(err)
	}

	handlerCalled := false
	err = c.parseProcfsFile(arcstatsFile, "arcstats", func(s zfsSysctl, v zfsMetricValue) {

		if s != zfsSysctl("kstat.zfs.misc.arcstats.hits") {
			return
		}

		handlerCalled = true

		if v != zfsMetricValue(8772612) {
			t.Fatalf("Incorrect value parsed from procfs data")
		}

	})

	if err != nil {
		t.Fatal(err)
	}

	if !handlerCalled {
		t.Fatal("Arcstats parsing handler was not called for some expected sysctls")
	}
}

func TestZfetchstatsParsing(t *testing.T) {
	zfetchstatsFile, err := os.Open("fixtures/proc/spl/kstat/zfs/zfetchstats")
	if err != nil {
		t.Fatal(err)
	}
	defer zfetchstatsFile.Close()

	c := zfsCollector{}
	if err != nil {
		t.Fatal(err)
	}

	handlerCalled := false
	err = c.parseProcfsFile(zfetchstatsFile, "zfetchstats", func(s zfsSysctl, v zfsMetricValue) {

		if s != zfsSysctl("kstat.zfs.misc.zfetchstats.hits") {
			return
		}

		handlerCalled = true

		if v != zfsMetricValue(7067992) {
			t.Fatalf("Incorrect value parsed from procfs data")
		}

	})

	if err != nil {
		t.Fatal(err)
	}

	if !handlerCalled {
		t.Fatal("Zfetchstats parsing handler was not called for some expected sysctls")
	}
}

func TestZilParsing(t *testing.T) {
	zilFile, err := os.Open("fixtures/proc/spl/kstat/zfs/zil")
	if err != nil {
		t.Fatal(err)
	}
	defer zilFile.Close()

	c := zfsCollector{}
	if err != nil {
		t.Fatal(err)
	}

	handlerCalled := false
	err = c.parseProcfsFile(zilFile, "zil", func(s zfsSysctl, v zfsMetricValue) {

		if s != zfsSysctl("kstat.zfs.misc.zil.zil_commit_count") {
			return
		}

		handlerCalled = true

		if v != zfsMetricValue(10) {
			t.Fatalf("Incorrect value parsed from procfs data")
		}

	})

	if err != nil {
		t.Fatal(err)
	}

	if !handlerCalled {
		t.Fatal("Zil parsing handler was not called for some expected sysctls")
	}
}

func TestVdevCacheStatsParsing(t *testing.T) {
	vdevCacheStatsFile, err := os.Open("fixtures/proc/spl/kstat/zfs/vdev_cache_stats")
	if err != nil {
		t.Fatal(err)
	}
	defer vdevCacheStatsFile.Close()

	c := zfsCollector{}
	if err != nil {
		t.Fatal(err)
	}

	handlerCalled := false
	err = c.parseProcfsFile(vdevCacheStatsFile, "vdev_cache_stats", func(s zfsSysctl, v zfsMetricValue) {

		if s != zfsSysctl("kstat.zfs.misc.vdev_cache_stats.delegations") {
			return
		}

		handlerCalled = true

		if v != zfsMetricValue(40) {
			t.Fatalf("Incorrect value parsed from procfs data")
		}

	})

	if err != nil {
		t.Fatal(err)
	}

	if !handlerCalled {
		t.Fatal("VdevCacheStats parsing handler was not called for some expected sysctls")
	}
}

func TestXuioStatsParsing(t *testing.T) {
	xuioStatsFile, err := os.Open("fixtures/proc/spl/kstat/zfs/xuio_stats")
	if err != nil {
		t.Fatal(err)
	}
	defer xuioStatsFile.Close()

	c := zfsCollector{}
	if err != nil {
		t.Fatal(err)
	}

	handlerCalled := false
	err = c.parseProcfsFile(xuioStatsFile, "xuio_stats", func(s zfsSysctl, v zfsMetricValue) {

		if s != zfsSysctl("kstat.zfs.misc.xuio_stats.onloan_read_buf") {
			return
		}

		handlerCalled = true

		if v != zfsMetricValue(32) {
			t.Fatalf("Incorrect value parsed from procfs data")
		}

	})

	if err != nil {
		t.Fatal(err)
	}

	if !handlerCalled {
		t.Fatal("XuioStats parsing handler was not called for some expected sysctls")
	}
}
