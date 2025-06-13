from abc import ABC, abstractmethod

class IDatabase(ABC):
    @abstractmethod
    def apply_migrations(self):
        raise NotImplementedError