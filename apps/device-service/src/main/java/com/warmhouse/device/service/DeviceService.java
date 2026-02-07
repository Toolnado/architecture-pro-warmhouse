package com.warmhouse.device.service;

import com.warmhouse.device.model.Device;
import com.warmhouse.device.repository.DeviceRepository;
import java.util.List;
import org.springframework.stereotype.Service;

@Service
public class DeviceService {

  private final DeviceRepository repository;

  public DeviceService(DeviceRepository repository) {
    this.repository = repository;
  }

  public Device registerDevice(Device device) {
    return repository.save(device);
  }

  public List<Device> getAllDevices() {
    return repository.findAll();
  }
}
